const dateFormatter = new Intl.DateTimeFormat('en-US', {
  dateStyle: 'full',
  timeStyle: 'short',
});

export function formatDate(input: string) {
  return dateFormatter.format(new Date(input));
}
