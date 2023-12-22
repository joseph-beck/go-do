'use client'

import Link from "next/link";

export default async function Page() {
  return (
    <main className="w-full h-screen flex flex-col items-center justify-center px-4">
      <div className="max-w-sm w-full text-gray-600">
        <div className="text-center">
          <div className="mt-5 space-y-2">
            <h3 className="text-indigo-400 text-2xl font-bold sm:text-3xl">log in to your account</h3>
            <p className="text-indigo-300">don&apos;t have an account? <Link href="/signup" className="font-medium text-indigo-600 hover:text-indigo-500">sign up &rarr;</Link></p>
          </div>
        </div>
        <form
          onSubmit={(e) => e.preventDefault()}
          className="mt-8 space-y-5"
        >
          <div>
            <div className="py-2">
              <label className="font-medium relative">
                email
              </label>
              <div className="relative group">
                <div className="absolute mt-2 -inset-0.5 bg-gradient-to-r from-pink-600 to-purple-600 rounded-lg blur opacity-75 group-hover:opacity-100 transition duration-1000 group-hover:duration-200 animate-tilt"></div>
                <input
                  type="email"
                  required
                  className="relative w-full mt-2 px-3 py-2 bg-black rounded-lg leading-0 text-indigo-400 hover:text-gray-200 transition duration-1000 group-hover:duration-200"
                />
              </div>
            </div>
            <div className="py-2">
              <label className="font-medium">
                password
              </label>
              <div className="relative group">
                <div className="absolute mt-2 -inset-0.5 bg-gradient-to-r from-pink-600 to-purple-600 rounded-lg blur opacity-75 group-hover:opacity-100 transition duration-1000 group-hover:duration-200 animate-tilt"></div>
                <input
                  type="password"
                  required
                  className="relative w-full mt-2 px-3 py-2 bg-black rounded-lg leading-0 text-indigo-400 hover:text-gray-200 transition duration-1000 group-hover:duration-200"
                />
              </div>
            </div>
          </div>

          <div className="py-2">
            <div className="relative group">
              <div className="absolute -inset-0.5 bg-gradient-to-r from-pink-600 to-purple-600 rounded-lg blur opacity-75 group-hover:opacity-100 transition duration-1000 group-hover:duration-200 animate-tilt"></div>
              <button
                className="relative bg-gradient-to-r from-indigo-500 to-pink-500 rounded-lg leading-0 w-full px-4 py-2 font-medium text-black hover:text-gray-200 transition duration-1000 group-hover:duration-200"
              >
                sign in
              </button>
            </div>

            <div className="text-center py-4">
              <Link href="/support" className="font-medium text-indigo-600 hover:text-indigo-500">forgot password?</Link>
            </div>
          </div>

        </form>
      </div>
    </main>
  );
};
