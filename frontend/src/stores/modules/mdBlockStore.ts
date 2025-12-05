import { defineStore } from "pinia";
import { ref, computed, watch } from "vue";
import * as textChunkService from "@/services/textChunkService";
import { MdBlock } from "@/types/model/mdBlock";

export const useMdBlockStore = defineStore("blockStore", () => {
  // array to hold markdown blocks
  const blocks = ref<MdBlock[]>([]);
  const loadingSplit = ref(false);

  // selected block index, no block selected when -1
  const selectedBlockIndex = ref<number>(-1);

  const hasBlocks = computed(() => blocks.value.length > 0);

  const selectedBlock = computed(() => {
    if (
      selectedBlockIndex.value < 0 ||
      selectedBlockIndex.value >= blocks.value.length
    )
      return null;
    return blocks.value[selectedBlockIndex.value] ?? null;
  });

  async function fetchChunks(
    text: string,
    strategy: string,
    params: Record<string, any> = {},
  ): Promise<MdBlock[] | null> {
    loadingSplit.value = true;
    try {
      const res = await textChunkService.splitText(text, strategy, params);
      if (!res) {
        blocks.value = [];
        selectedBlockIndex.value = -1;
        return null;
      }

      blocks.value = res;
      selectedBlockIndex.value = res.length > 0 ? 0 : -1;
      return res;
    } catch (err) {
      console.error("fetchChunks failed", err);
      blocks.value = [];
      selectedBlockIndex.value = -1;
      return null;
    } finally {
      loadingSplit.value = false;
    }
  }

  function setBlocks(newBlocks: MdBlock[]) {
    blocks.value = newBlocks;
    selectedBlockIndex.value = blocks.value.length > 0 ? 0 : -1;
  }

  function insertBlock(block: MdBlock, index: number) {
    if (index < 0 || index >= blocks.value.length) {
      blocks.value.push(block);
    } else {
      blocks.value.splice(index, 0, block);
    }
    // select the newly inserted block
    selectedBlockIndex.value = blocks.value.indexOf(block);
  }

  function removeBlock(index: number): boolean {
    if (index < 0 || index >= blocks.value.length) return false;

    // remove the block from the array
    blocks.value.splice(index, 1);

    // no blocks left -> no selection
    if (blocks.value.length === 0) {
      selectedBlockIndex.value = -1;
      return true;
    }

    // if nothing was selected before, nothing to adjust
    if (selectedBlockIndex.value < 0) {
      return true;
    }

    // if the removed block is before the selected block, shift selection left
    if (selectedBlockIndex.value > index) {
      selectedBlockIndex.value -= 1;
      return true;
    }

    // if the removed block is the selected block
    if (selectedBlockIndex.value === index) {
      // prefer previous block if exists, otherwise select the new block at current index
      selectedBlockIndex.value = index - 1 >= 0 ? index - 1 : 0;
      return true;
    }

    // if the removed block is after the selected block, selection remains unchanged
    return true;
  }

  function mergeBlock(nextIndex: number): boolean {
    if (nextIndex <= 0 || nextIndex >= blocks.value.length) return false;

    const prev = blocks.value[nextIndex - 1];
    const curr = blocks.value[nextIndex];

    // Merge text content
    const merged: MdBlock = {
      ...prev,
      content: (prev.content ?? "") + (curr.content ?? ""),
    };

    // Replace previous and remove current
    blocks.value.splice(nextIndex - 1, 2, merged);

    // Select merged block
    selectedBlockIndex.value = nextIndex - 1;

    return true;
  }

  function moveBlock(from: number, to: number): boolean {
    if (
      from < 0 ||
      from >= blocks.value.length ||
      to < 0 ||
      to >= blocks.value.length ||
      from === to
    ) {
      return false;
    }

    const [item] = blocks.value.splice(from, 1);
    blocks.value.splice(to, 0, item);

    // Keep selection on moved element
    selectedBlockIndex.value = to;
    return true;
  }

  function splitBlock(index: number, pos: number) {
    if (index < 0 || index >= blocks.value.length) return false;

    const original = blocks.value[index];
    const content = original.content ?? "";

    const left = content.slice(0, pos);
    const right = content.slice(pos);

    const leftBlock: MdBlock = {
      ...original,
      content: left,
    };

    const rightBlock: MdBlock = {
      id: crypto.randomUUID(),
      type: original.type,
      content: right,
    };

    blocks.value.splice(index, 1, leftBlock, rightBlock);

    selectedBlockIndex.value = index + 1;
    return true;
  }

  function clearBlocks() {
    blocks.value = [];
    selectedBlockIndex.value = -1;
  }

  return {
    blocks,
    loadingSplit,
    hasBlocks,
    selectedBlockIndex,
    selectedBlock,
    fetchChunks,
    setBlocks,
    insertBlock,
    removeBlock,
    splitBlock,
    moveBlock,
    mergeBlock,
    clearBlocks,
  };
});
