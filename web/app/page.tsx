import { NavigationBar } from '@/components/navigation/navbar';
import Link from 'next/link';

export default function Page() {
  const content = (
    <main>
      <NavigationBar />
      <Link href="wall">
        ToDo
      </Link>
    </main>
  );

  return content;
}
