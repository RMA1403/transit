import clsx from "clsx";
import localFont from "next/font/local";

import "./globals.css";
import { Inter, Poppins } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });

const poppinsLight = Poppins({
  subsets: ["latin"],
  variable: "--font-poppins-light",
  weight: "400",
});

const poppinsBold = Poppins({
  subsets: ["latin"],
  variable: "--font-poppins-bold",
  weight: "800",
});

const jeko = localFont({
  src: "../assets/fonts/FontsFree-Net-Jeko-Variable.ttf",
  variable: "--font-jeko",
});

const allrounderBook = localFont({
  src: "../assets/fonts/AllrounderGroteskTest-Book.otf",
  variable: "--font-allrounder-book",
});

const allrounderRegular = localFont({
  src: "../assets/fonts/AllrounderGroteskTest-Regular.otf",
  variable: "--font-allrounder-regular",
});

export const metadata = {
  title: "TransIT",
  description: "Solusi Distrupsi Moda Transportasi",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body
        className={clsx(
          "bg-slate-200",
          inter.className,
          jeko.variable,
          allrounderBook.variable,
          allrounderRegular.variable,
          poppinsLight.variable,
          poppinsBold.variable
        )}
      >
        <div className="max-w-[420px] bg-white mx-auto min-h-[100vh] overflow-hidden">
          {children}
        </div>
      </body>
    </html>
  );
}
