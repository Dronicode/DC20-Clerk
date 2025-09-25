import { useEffect, useState } from 'react';
import { getMe, subscribeMe } from './meStore';

export const useMe = () => {
  const [me, setMeState] = useState(getMe());

  useEffect(() => {
    const unsubscribe = subscribeMe(setMeState);
    return () => unsubscribe();
  }, []);

  return me;
};
