import { SpecialBlink } from '@/components/input/blink';
import { Footer } from '@/components/navigation/footer';
import { NavigationBar } from '@/components/navigation/navbar';

export default function Page() {
  const content = (
    <main>
      <div className="flex flex-col h-screen">
        <NavigationBar />

        <div className="flex-grow min-h-full">
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
