import { Divider } from "../layout/divider";
import { Footer } from "../navigation/footer";

export const Page303 = (): JSX.Element => {
  const content = (
    <div>
      <div className="flex flex-col h-screen">
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
              <a href="/" className="text-indigo-600 duration-200 bg-black rounded-lg leading-none hover:text-indigo-400 font-medium inline-flex items-center gap-x-1">
                go home &rarr;
              </a>
            </div>
          </div>
        </div>
      </div>

      <Footer />
    </div>
  );

  return content;
};
