type StorageCallback = (
  key: string,
  oldValue: string | null,
  newValue: string | null,
) => void;

/**
 * Subscribes to localStorage changes across tabs.
 * Only triggers in tabs other than the one that made the change.
 */
export const onStorageChange = (callback: StorageCallback) => {
  const handler = (e: StorageEvent) => {
    if (e.storageArea !== localStorage) return;
    callback(e.key ?? '', e.oldValue, e.newValue);
  };

  window.addEventListener('storage', handler);
  return () => window.removeEventListener('storage', handler);
};
