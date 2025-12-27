<script setup lang="ts">
import { ref, nextTick, onUnmounted } from "vue";
import { useMdBlockStore } from "@/stores";
import draggable from "vuedraggable";
import { marked } from "marked";
import { NDropdown, useDialog } from "naive-ui";
import { useNotification } from "@/composables/useNotification";

// Get the markdown block store for managing text blocks
const blockStore = useMdBlockStore();
const blocks = blockStore.blocks;
const { showSuccess } = useNotification();

// Initialize dialog for unified confirmation dialogs
const dialog = useDialog();

// Store references to textarea elements for cursor manipulation and auto-sizing
const textAreaRefs = ref<(HTMLTextAreaElement | null)[]>([]);

// Set textarea reference and adjust its height to fit content
function setTextAreaRef(index: number, el: any) {
  if (el) {
    textAreaRefs.value[index] = el;
    adjustTextAreaHeight(el);
  }
}

// Auto-adjust textarea height to match scrollHeight for dynamic content
function adjustTextAreaHeight(textarea: HTMLTextAreaElement) {
  textarea.style.height = "auto";
  textarea.style.height = textarea.scrollHeight + "px";
}

// Parse markdown content and render it as HTML for preview
function renderMarkdown(text: string) {
  const html = marked.parseInline(text ?? "");
  if (html instanceof Promise) {
    return html.then((h) => h.replace(/style="[^"]*"/g, ""));
  }
  return html;
}

// Handle keyboard events in the editor textarea (Enter, Delete, etc.)
function onKeyDown(event: KeyboardEvent, index: number) {
  const block = blocks[index];
  const textarea = textAreaRefs.value[index];
  if (!textarea) return;
  const pos = textarea.selectionStart ?? 0;

  // Handle Enter key: create new block or split current block
  if (event.key === "Enter") {
    event.preventDefault();
    if (pos === (block.content?.length ?? 0)) {
      // Cursor at end of block - create new block
      const newBlock = {
        id: crypto.randomUUID(),
        type: block.type,
        content: "",
      };
      blockStore.insertBlock(newBlock, index + 1);
      nextTick(() => {
        textAreaRefs.value[index + 1]?.focus();
      });
    } else {
      // Cursor in middle of block - split into two blocks
      blockStore.splitBlock(index, pos);
      nextTick(() => {
        textAreaRefs.value[index + 1]?.focus();
      });
    }
  } else if (event.key === "Delete" && pos === 0) {
    // Handle Delete key at beginning of block: delete block or merge with previous
    event.preventDefault();
    if (index === 0) {
      // First line - delete block
      blockStore.removeBlock(index);
      nextTick(() => {
        textAreaRefs.value[0]?.focus();
      });
    } else {
      // Merge with previous block
      blockStore.mergeBlock(index);
      nextTick(() => {
        const prevLength = blocks[index - 1]?.content?.length ?? 0;
        const prevTextarea = textAreaRefs.value[index - 1];
        prevTextarea?.focus();
        prevTextarea?.setSelectionRange(prevLength, prevLength);
      });
    }
  }
}

// Context menu configuration and state management
const contextMenuOptions = ref<Array<{ label: string; key: string }>>([]);
const contextMenuIndex = ref<number>(-1);

// Context menu display state and position
const contextMenu = ref<{
  show: boolean;
  pos: { x: number; y: number };
}>({
  show: false,
  pos: { x: 0, y: 0 },
});

// Display context menu at mouse position with block actions
function showContextMenu(event: MouseEvent, index: number) {
  contextMenuIndex.value = index;
  contextMenu.value.pos = { x: event.clientX, y: event.clientY };
  contextMenuOptions.value = [
    { label: "Delete", key: "delete" },
    { label: "Insert Above", key: "insert-above" },
    { label: "Insert Below", key: "insert-below" },
  ];
}

// Handle context menu item selection
function handleMenuSelect(key: string) {
  const index = contextMenuIndex.value;
  if (index === -1) return;

  switch (key) {
    case "delete":
      blockStore.removeBlock(index);
      break;
    case "insert-above":
      blockStore.insertBlock(
        { id: crypto.randomUUID(), type: blocks[index].type, content: "" },
        index,
      );
      break;
    case "insert-below":
      blockStore.insertBlock(
        { id: crypto.randomUUID(), type: blocks[index].type, content: "" },
        index + 1,
      );
      break;
  }
}

// Clear all blocks with confirmation dialog
function clearAllBlocks() {
  if (blocks.length === 0) return;
  dialog.warning({
    title: "Clear All Blocks",
    content: `Are you sure you want to clear all ${blocks.length} block${blocks.length !== 1 ? "s" : ""}? This action cannot be undone.`,
    positiveText: "Clear",
    negativeText: "Cancel",
    onPositiveClick: () => {
      blockStore.clearBlocks();
      showSuccess(
        `Cleared ${blocks.length} block${blocks.length !== 1 ? "s" : ""}`,
      );
    },
  });
}

// Close context menu when clicking elsewhere on the document
const handleDocumentClick = () => {
  contextMenu.value.show = false;
};

// Cleanup event listener on component unmount
onUnmounted(() => {
  document.removeEventListener("click", handleDocumentClick);
});
</script>

<template>
  <!-- Context menu dropdown for block operations -->
  <n-dropdown
    placement="top-start"
    trigger="manual"
    :x="contextMenu.pos.x"
    :y="contextMenu.pos.y"
    :options="contextMenuOptions"
    :show="contextMenu.show"
    @clickoutside="contextMenu.show = false"
    @select="handleMenuSelect"
  />
  <!-- Header with block counter and clear button -->
  <div class="editor-header">
    <span class="block-counter"
      >{{ blocks.length }} block{{ blocks.length !== 1 ? "s" : "" }}</span
    >
    <button
      v-if="blocks.length > 0"
      class="clear-btn"
      @click="clearAllBlocks"
      title="Clear all blocks"
    >
      Clear All
    </button>
  </div>
  <!-- Main editor table with split-pane layout -->
  <table class="block-table">
    <!-- Draggable list of markdown blocks -->
    <draggable v-model="blocks" handle=".drag-handle" item-key="id" tag="tbody">
      <template #item="{ element: block, index }">
        <tr
          @contextmenu.prevent="
            showContextMenu($event, index);
            contextMenu.show = true;
          "
          class="draggable-item"
        >
          <!-- Left column: editable textarea for text input -->
          <td class="text-pane">
            <textarea
              :ref="(el) => setTextAreaRef(index, el)"
              v-model="block.content"
              @keydown="onKeyDown($event, index)"
              @input="
                adjustTextAreaHeight($event.target as HTMLTextAreaElement)
              "
              class="editor-textarea"
            />
            <!-- Drag handle icon -->
            <span class="drag-handle">⠿</span>
          </td>

          <!-- Right column: live markdown preview/rendering -->
          <td class="markdown-pane" v-html="renderMarkdown(block.content)"></td>
        </tr>
      </template>
    </draggable>
  </table>
</template>

<style scoped>
/* Editor header with block counter and controls */
.editor-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px;
  background-color: #f5f5f5;
  border-bottom: 1px solid #e0e0e0;
  margin-bottom: 8px;
}

/* Block counter display */
.block-counter {
  font-size: 14px;
  font-weight: 500;
  color: #333;
}

/* Clear all blocks button */
.clear-btn {
  padding: 6px 12px;
  background-color: #ff6b6b;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 13px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.clear-btn:hover {
  background-color: #ff5252;
}

.clear-btn:active {
  background-color: #e64545;
}

/* Main editor table styling */
.block-table {
  width: 100%;
  border-collapse: collapse;
}

/* Table row displays as flex for split-pane layout */
.block-table tr {
  display: flex;
}

/* Base table cell styling */
.block-table td {
  vertical-align: top;
  padding: 12px;
  display: flex;
  flex-direction: column;
}

/* Left column: text pane (50% width) */
.text-pane {
  width: 50%;
  position: relative;
}

/* Textarea styling: auto-expanding, transparent background */
.editor-textarea {
  width: 100%;
  resize: none;
  border: none;
  font-family: inherit;
  font-size: 14px;
  padding: 0;
  margin: 0;
  background-color: transparent;
  line-height: 1.5;
  overflow: hidden;
  flex: 1;
}

/* Remove default textarea focus outline */
.editor-textarea:focus {
  outline: none;
}

/* Drag handle styling */
.drag-handle {
  cursor: grab;
  position: absolute;
  top: 8px;
  right: 8px;
  user-select: none;
}

/* Visual feedback when dragging */
.drag-handle:active {
  cursor: grabbing;
}

/* Right column: markdown pane (50% width) */
.markdown-pane {
  width: 50%;
  overflow: auto;
  font-size: 14px;
  line-height: 1.5;
  flex: 1;
}

/* Remove default paragraph margins in markdown preview */
.markdown-pane :deep(p) {
  margin: 0;
}
</style>
