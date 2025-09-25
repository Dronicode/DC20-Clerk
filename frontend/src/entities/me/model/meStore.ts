type MeState = {
  accessToken: string | null;
  email: string | null;
};

let me: MeState = {
  accessToken: null,
  email: null,
};

let listeners: ((state: MeState) => void)[] = [];

export const setMe = (data: Partial<MeState>) => {
  me = { ...me, ...data };
  listeners.forEach((fn) => fn(me));
};

export const clearMe = () => {
  me = { accessToken: null, email: null };
  listeners.forEach((fn) => fn(me));
};

export const getMe = () => me;

export const subscribeMe = (fn: (state: MeState) => void) => {
  listeners.push(fn);
  return () => {
    listeners = listeners.filter((l) => l !== fn);
  };
};
