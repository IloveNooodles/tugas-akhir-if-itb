export function generateColumnsFromArray<T>(arr: T[], filter?: string[]) {
  const result: {
    [key: string]: any;
    key: string;
    sortable?: boolean | undefined;
    sort?: ((a: any, b: any, direction: 'asc' | 'desc') => number) | undefined;
    direction?: 'asc' | 'desc' | undefined;
    class?: string | undefined;
  }[] = [];

  if (!arr || arr.length === 0) {
    return [];
  }

  const firstEle = arr[0];
  let res = Object.entries(firstEle).map(([k, v]) => {
    const transformed: Column = {
      key: k,
      label: k,
      sortable: true,
    };

    return transformed;
  });

  res.push({ key: 'actions' });
  res = res.filter((item) => !filter.includes(item.key));
  return res;
}

export function genereateSelectFromArray(
  arr: any[],
  display: string,
  value: string,
  filter?: any[],
) {
  let result = arr;
  if (arr.length === 0) {
    return result;
  }

  const filterId = filter.map((v) => v.id)
  console.log(filterId)
  if(filterId.length >= 0){
    result = arr.filter((v) => !filterId.includes(v.id));
  }

  result = result.map((v) => {
    return {
      name: v[display],
      value: v[value],
    };
  });

  return result;
}
