export interface BaseCardProps {
  titel?: string;
  description?: string;

  info?: string;
  infoTitel?: string;

  imageUrl?: string;
  allowPlacehoderImage?: boolean;

  href?: string;
}

/**
 * This is the Base card used to construct cards like transaction or group
 * @param props {@link BaseCardProps}
 * @returns
 */
export function BaseCard(
  props: BaseCardProps,
) {
  const {
    titel,
    description,
    info,
    infoTitel,
    imageUrl,
    allowPlacehoderImage = true,
    href,
  } = props;

  return (
    <a
      class={`p-4 rounded-xl bg-primary flex items-center gap-6 ${
        href && "cursor-pointer"
      }`}
      href={href}
    >
      {/* Base card image */}
      {imageUrl && (
        <img
          src={imageUrl}
          class="rounded-xl aspect-square"
        />
      )}
      {allowPlacehoderImage && !imageUrl && (
        <div class="rounded-xl aspect-square bg-neutral h-16" />
      )}

      {/* Base card content */}
      <div class="flex-1 h-16 overflow-hidden
        flex flex-col justify-between
        text-primary-content">
        <span class="max-w-full text-xl font-bold overflow-hidden whitespace-nowrap text-ellipsis">
          {titel}
        </span>
        <div class="max-w-full overflow-hidden whitespace-nowrap text-ellipsis
          brightness-75">
          {description}
        </div>
      </div>

      {/* Card info only visible if set*/}
      {(info || infoTitel) && (
        <div class="ml-auto h-16 overflow-hidden
        flex flex-col justify-between
        text-right text-primary-content">
          <span class="max-w-full text-xl font-bold whitespace-nowrap">
            {info}
          </span>
          <span class="text-sm text-primary-content brightness-75">
            {infoTitel}
          </span>
        </div>
      )}
    </a>
  );
}
