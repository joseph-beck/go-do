import { Page404 } from "@/components/status/404";

export default async function Page() {
  const content = (
    <main>
      <Page404 />
    </main>
  );

  return content;
}
