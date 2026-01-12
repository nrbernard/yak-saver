import type { ButtonHTMLAttributes } from "react";

export function Button(
  props: Omit<ButtonHTMLAttributes<HTMLButtonElement>, "className">
) {
  const { children, ...rest } = props;
  return (
    <button
      className={`px-4 py-1 text-sm hover:cursor-pointer rounded-md bg-banana hover:text-white dark:bg-strawberry dark:text-dark-text`}
      {...rest}
    >
      {children}
    </button>
  );
}
