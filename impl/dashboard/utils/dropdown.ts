import type { DropdownItem } from '#ui/types/dropdown';

export const dropdownItems: DropdownItem[][] = [
  [
    {
      label: 'Detail',
      icon: 'i-heroicons-document-magnifying-glass-16-solid',
      click: (e) => console.log(e)
    },
    {
      label: 'Edit',
      icon: 'i-heroicons-pencil-square-20-solid',
    },
  ],
  [
    {
      label: 'Delete',
      icon: 'i-heroicons-trash-20-solid',
    },
  ],
];
