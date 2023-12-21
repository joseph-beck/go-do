import { Divider } from "@/components/layout/divider";
import { NavigationBar } from "@/components/navigation/navbar";

export default async function Page() {
  const content = (
    <main>
      <NavigationBar />

      <div className="max-w-screen-xl mx-auto px-4 flex items-center justify-start h-screen md:px-8">
        <div className="max-w-lg mx-auto space-y-3 text-center">
          <h3 className="text-gray-200 text-4xl font-semibold sm:text-5xl">
            Page not found
          </h3>
          <Divider width="1/4" />
          <p className="text-gray-400">
            Sorry, the page you are looking for could not be found or has been removed.
          </p>
          <a href="/" className="text-indigo-600 duration-200 bg-black rounded-lg leading-none hover:text-indigo-400 font-medium inline-flex items-center gap-x-1">
            Go home &rarr;
          </a>
        </div>
      </div>
    </main>
  );

  return content;
}
