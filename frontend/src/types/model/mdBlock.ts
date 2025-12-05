/**
 * BlockType — categories of editor blocks.
 *
 * - `paragraph`: meaningful markdown content that should be persisted
 *   and sent to the backend.
 *
 * - `meta`: auxiliary blocks used only for UI or structural/annotation purposes.
 *   Meta blocks MAY contain text (e.g., notes, tips, UI hints), but their content
 *   is not considered part of the real document and is excluded from backend payloads.
 */
export type BlockType = "paragraph" | "meta";

/**
 * Block — a single editable, structural, or display block in the editor.
 *
 * @property id - Unique identifier used for ordering and operations (insert/delete).
 * @property type - Semantic category: `'paragraph'` or `'meta'`.
 * @property content - Original markdown or text; may be filled for meta blocks
 *                     but will not be submitted to backend.
 */
export interface MdBlock {
  id: string;
  type: BlockType;
  content: string;
}
