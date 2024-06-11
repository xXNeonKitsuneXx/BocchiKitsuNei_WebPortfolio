import '../../app/globals.css';
import { Contact } from '@/components/contactpage/contact';
import { Footer } from '@/components/footer/footer';
import { Toaster } from "@/components/ui/toaster";


export const ContactPage = () => {

  return (
    <>
    <Toaster />
    <Contact />
    <Footer />
    </>
  );
};
