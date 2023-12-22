import { SpecialBlink } from '@/components/input/blink';
import { Footer } from '@/components/navigation/footer';
import { NavigationBar } from '@/components/navigation/navbar';

export default function Page() {
  const content = (
    <main>
      <div className="flex flex-col">
        <NavigationBar />

        <div className="flex-grow h-screen">
          <div className="mt-16 px-4 justify-center text-center md:px-8">
            <h3 className="text-indigo-400 text-3xl mb-8">go and get yourself organised today</h3>
            <p className="text-indigo-300">store all your tasks in one place, godo!</p>
          </div>
          <div className="max-w-screen-xl mx-auto px-4 flex items-center justify-center h-full md:px-8">
            <div className="max-w-lg mx-auto space-y-3 text-center">
              <SpecialBlink href="wall" text="godo" />
            </div>
          </div>
        </div>

        <Footer />
      </div>
    </main>
  );

  return content;
}
