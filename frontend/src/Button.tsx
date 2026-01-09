import type { ButtonHTMLAttributes } from "react";

export function Button(props: ButtonHTMLAttributes<HTMLButtonElement>) {
  const { children, className, ...rest } = props;
  return (
    <button
      className={`px-4 py-1 text-sm hover:cursor-pointer rounded-md bg-banana hover:text-white ${className}`}
      {...rest}
    >
      {children}
    </button>
  );
}
