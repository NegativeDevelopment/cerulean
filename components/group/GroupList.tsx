import { GroupListEntry } from "./GroupListEntry.tsx"

interface GroupListProps {
  groups: {name: string}[]
}

export function GroupList({groups}: GroupListProps) {
  return (
    <div>
      {groups.map(group => <GroupListEntry group={group}/>)}
    </div>
  )
}
