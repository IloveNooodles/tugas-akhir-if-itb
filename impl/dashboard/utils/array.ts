export function generateColumnsFromArray<T>(arr: T[]) {
  const result: {
    [key: string]: any;
    key: string;
    sortable?: boolean | undefined;
    sort?: ((a: any, b: any, direction: 'asc' | 'desc') => number) | undefined;
    direction?: 'asc' | 'desc' | undefined;
    class?: string | undefined;
  }[] = [];

  if (arr.length === 0) {
    return result;
  }

  const firstEle = arr[0];
  const res = Object.entries(firstEle).map(([k, v]) => {
    const transformed: Column = {
      key: k,
      label: k,
      sortable: true,
    };

    return transformed;
  });

  res.push({ key: 'actions' });
  return res;
}
