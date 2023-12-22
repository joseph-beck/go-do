import { Blink, SpecialBlink } from "../input/blink";
import { Divider } from "../layout/divider";

type Navs = {
  title: string;
  path: string;
};

export const NavigationBar = (): JSX.Element => {
  const navigation: Navs[] = [
    { title: "home", path: "/" },
    { title: "features", path: "/features" },
    { title: "help", path: "/help" },
  ]

  const content = (
    <nav className="bg-black w-full border-b md:border-0 md:static">
      <div className="items-center px-4 max-w-screen-xl p-8 mx-auto md:flex md:px-8">
        <div className="flex-1 justify-self-center pb-3 mt-8 md:block md:pb-0 md:mt-0 block">
          <ul className="justify-center items-center space-y-8 md:flex md:space-x-12 md:space-y-0">
            {
              navigation.map((item, idx) => {
                return (
                  <div key={idx}>
                    <Blink text={item.title} href={item.path} />
                  </div>
                );
              })
            }
            <div className="pl-32">
              <SpecialBlink text="login &rarr;" href="/login" />
            </div>
          </ul>
        </div>
      </div>

      <Divider width="3/4" />
    </nav>
  );

  return content;
};
