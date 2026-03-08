import { useState, useEffect } from "react";

export const SERVER_BASE_URL = import.meta.env.VITE_SERVER_BASE_URL;
export const API_BASE_URL_V1 = import.meta.env.VITE_API_BASE_URL;

export const useLocalStorage = (key, defaultValue) => {
  const [value, setValue] = useState(() => {
    const currentValue = localStorage.getItem(key);

    if (currentValue) {
      return JSON.parse(currentValue);
    } else {
      return defaultValue;
    }
  });

  useEffect(() => {
    localStorage.setItem(key, JSON.stringify(value));
  }, [value]);

  return [value, setValue];
};
