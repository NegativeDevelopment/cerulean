import { Group } from "../../types/group.ts";
import { GroupListEntry } from "./GroupListEntry.tsx"

interface GroupListProps {
  groups: Group[]
}

export function GroupList({groups}: GroupListProps) {
  return (
    <div>
      {groups.map(group => <GroupListEntry group={group} key={group.id}/>)}
    </div>
  )
}
