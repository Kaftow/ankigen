import {
  GetSupportedExtensions,
  SelectFile,
  ExtractText,
} from "../../wailsjs/go/main/App";

export async function getSupportedExtensions(): Promise<string[]> {
  try {
    const res = (await GetSupportedExtensions()) as unknown;
    return Array.isArray(res) ? (res as string[]) : [];
  } catch (err) {
    console.error("getSupportedExtensions failed", err);
    return [];
  }
}

export async function selectFile(): Promise<string[] | null> {
  try {
    const res = (await SelectFile()) as unknown;
    if (Array.isArray(res)) return res as string[];
    if (typeof res === "string" && res) return [res];
    return null;
  } catch (err) {
    console.error("selectFile failed", err);
    return null;
  }
}

export async function extractText(path: string): Promise<string | null> {
  try {
    const res = (await ExtractText(path)) as unknown;
    return typeof res === "string" ? res : null;
  } catch (err) {
    console.error("extractText failed", err);
    return null;
  }
}

export async function pickAndRead(
  allowedExtensions?: string[],
): Promise<{ filename: string; content: string } | null> {
  const paths = await selectFile();
  if (!paths || paths.length === 0) return null;
  const fullPath = paths[0];

  // Extract filename from path (works for both / and \ separators)
  const parts = fullPath.split(/[\\/]/);
  const filename = parts[parts.length - 1] || fullPath;

  // If allowedExtensions provided, validate extension before extracting
  if (Array.isArray(allowedExtensions) && allowedExtensions.length > 0) {
    const dotIndex = filename.lastIndexOf(".");
    const ext = dotIndex >= 0 ? filename.slice(dotIndex + 1).toLowerCase() : "";
    const normalized = allowedExtensions.map((s) =>
      (s.startsWith(".") ? s.slice(1) : s).toLowerCase(),
    );
    if (!ext) {
      console.warn(
        "pickAndRead: file has no extension, skipping extraction:",
        filename,
      );
      return null;
    }
    if (!normalized.includes(ext)) {
      console.warn(
        `pickAndRead: unsupported file extension .${ext}, allowed: ${allowedExtensions.join(", ")}`,
      );
      return null;
    }
  }

  const content = await extractText(fullPath);
  if (content == null) return null;
  return { filename, content };
}
