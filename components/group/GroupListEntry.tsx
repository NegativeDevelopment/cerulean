import { Group } from "../../types/group.ts"
import { IconChevronRight } from "./../icons/IconChevronRight.tsx"

interface GroupListEntryProps {
  group: Group
}

export function GroupListEntry({group}: GroupListEntryProps) {
  return (
    <div class="m-2 p-2 rounded bg-gray-400 text-gray-50 font-bold flex justify-between h-12 items-center">
      <div>{ group.name }</div>
      <a href={`/groups/${group.id}`}>
        <IconChevronRight />
      </a>
    </div>
  )
}
