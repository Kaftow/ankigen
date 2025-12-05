import { SplitText } from "../../wailsjs/go/api/ChunkerAPI";
import { Chunk, ChunkConfig } from "@/types/api/textChunk";
import { MdBlock } from "@/types/model/mdBlock";

function chunkToMdBlock(chunk: Chunk): MdBlock {
  return {
    id: chunk.ID,
    type: "paragraph",
    content: chunk.Text,
  };
}

export async function splitText(
  text: string,
  strategy: string,
  params: Record<string, any> = {},
): Promise<MdBlock[] | null> {
  try {
    const config: ChunkConfig = { Strategy: strategy, Params: params };

    const res = await SplitText(text, config);
    if (!Array.isArray(res)) return null;

    return res.map((chunk) => chunkToMdBlock(chunk));
  } catch (err) {
    console.error(`splitText failed (strategy=${strategy})`, err);
    return null;
  }
}
