import { useMessage, useDialog } from "naive-ui";

export function useNotification() {
  const message = useMessage();
  const dialog = useDialog();

  // Modal error dialog
  function showError(content: string) {
    return new Promise<void>((resolve) => {
      dialog.error({
        title: "Error",
        content,
        positiveText: "OK",
        maskClosable: false,
        closeOnEsc: false,
        onPositiveClick: () => {
          resolve();
        },
      });
    });
  }

  // Modal error dialog with custom title
  function showErrorDialog(title: string, content: string) {
    return new Promise<void>((resolve) => {
      dialog.error({
        title,
        content,
        positiveText: "OK",
        maskClosable: false,
        closeOnEsc: false,
        onPositiveClick: () => {
          resolve();
        },
      });
    });
  }

  // Success notification (non-modal)
  function showSuccess(content: string) {
    message.success(content);
  }

  // Warning notification (non-modal)
  function showWarning(content: string) {
    message.warning(content);
  }

  // Info notification (non-modal)
  function showInfo(content: string) {
    message.info(content);
  }

  return {
    showError,
    showErrorDialog,
    showSuccess,
    showWarning,
    showInfo,
  };
}
