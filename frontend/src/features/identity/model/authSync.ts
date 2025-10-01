let setLoginState: ((loggedIn: boolean) => void) | null = null;

export const registerLoginStateSetter = (fn: (loggedIn: boolean) => void) => {
  setLoginState = fn;
};

export const updateLoginState = (loggedIn: boolean) => {
  if (setLoginState) setLoginState(loggedIn);
};
