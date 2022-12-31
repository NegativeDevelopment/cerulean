interface ProfilePictureProps {
  username: string;
  imageUrl?: string;
}

export default function ProfilePicture(
  { imageUrl, username }: ProfilePictureProps,
) {
  return (
    <></>
    //   {imageUrl && <img></img>}
    //   {!imageUrl && <div class="flex justify-center items-center rounded-full text-2xl text-green-700 bg-green-300 w-16">
    //   M
    // </div>}
  );
}
