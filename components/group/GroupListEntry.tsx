import { IconChevronRight } from "./../icons/IconChevronRight.tsx"

interface GroupListEntryProps {
  group: {
    name: string
  }
}

export function GroupListEntry({group}: GroupListEntryProps) {
  return (
    <div class="m-2 p-2 rounded bg-yellow-200 flex justify-between h-12 items-center">
      <div>{ group.name }</div>
      <IconChevronRight />
    </div>
  )
}
