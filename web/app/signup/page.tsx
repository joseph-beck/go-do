import { Divider } from "@/components/layout/divider";
import { Footer } from "@/components/navigation/footer";
import { NavigationBar } from "@/components/navigation/navbar";
import Link from "next/link";

export default async function Page() {
  const content = (
    <main>
      <div className="flex flex-col h-screen">
        <NavigationBar />

        <div className="flex-grow">
          <div className="max-w-screen-xl mx-auto px-4 flex items-center justify-center h-full md:px-8">
            <div className="max-w-lg mx-auto space-y-3 text-center">
              <h3 className="text-gray-200 text-4xl font-semibold sm:text-5xl">
                page under construction
              </h3>
              <Divider width="1/4" />
              <p className="text-gray-400">
                sorry, the page you are looking for is currently under construction.
              </p>
              <Link href="/" className="text-indigo-600 duration-200 bg-black rounded-lg leading-none hover:text-indigo-400 font-medium inline-flex items-center gap-x-1">
                go home &rarr;
              </Link>
            </div>
          </div>
        </div>

        <Footer />
      </div>
    </main>
  );

  return content;
}
