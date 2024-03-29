import { Divider } from "../layout/divider";

export const Page404 = (): JSX.Element => {
  return (
    <div className="flex flex-col h-screen">
      <main className="flex-grow">
        <div className="max-w-screen-xl mx-auto px-4 flex items-center justify-center h-full md:px-8">
          <div className="max-w-lg mx-auto space-y-3 text-center">
            <h3 className="text-gray-200 text-4xl font-semibold sm:text-5xl">
              forbidden
            </h3>
            <Divider width="1/4" />
            <p className="text-gray-400">
              sorry, you do not have access to this, if you think this is a mistake please try again.
            </p>
            <a href="/" className="text-indigo-600 duration-200 bg-black rounded-lg leading-none hover:text-indigo-400 font-medium inline-flex items-center gap-x-1">
              go home &rarr;
            </a>
          </div>
        </div>
      </main>
    </div>
  );
};
