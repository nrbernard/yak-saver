import type { ButtonHTMLAttributes } from "react";

export function Button(props: ButtonHTMLAttributes<HTMLButtonElement>) {
  const { children, className, ...rest } = props;
  return (
    <button
      className={`px-4 py-1 text-sm hover:cursor-pointer rounded-md outline-1 outline-banana text-banana hover:bg-banana hover:text-tangerine hover:outline-tangerine ${className}`}
      {...rest}
    >
      {children}
    </button>
  );
}
