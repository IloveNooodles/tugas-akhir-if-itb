interface Column {
  [key: string]: any;
  key: string;
  sortable?: boolean | undefined;
  sort?: ((a: any, b: any, direction: 'asc' | 'desc') => number) | undefined;
  direction?: 'asc' | 'desc' | undefined;
  class?: string | undefined;
}
