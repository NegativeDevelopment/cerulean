import { Group } from "../../types/group.ts";
import { BaseCard } from "./BaseCard.tsx";

export interface GroupCardProps {
  group: Group;
}

export function GroupCard(
  { group }: GroupCardProps,
) {
  return (
    <BaseCard
      titel={group.name}
      description={`ID: ${group.id}`}
      href={`/groups/${group.id}`}
    />
  );
}
