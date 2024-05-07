import type { DropdownItem } from '#ui/types/dropdown';

type Options = {
  to?: string;
  editFunc?: Function;
  deleteFunc?: Function;
};

export type GenerateDropdownItems = (row: any) => DropdownItem[][];
