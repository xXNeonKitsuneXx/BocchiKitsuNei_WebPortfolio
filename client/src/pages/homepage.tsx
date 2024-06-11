import '../../app/globals.css';
import { Hero } from '@/components/homepage/hero';
import { About } from '@/components/homepage/about';
import { Skill } from '@/components/homepage/skill';
import { HomeProject } from '@/components/homepage/project';
import { Footer } from '@/components/footer/footer';

export const HomePage = () => {

  return (
    <>

    <Hero />
    <About />
    <Skill />
    <HomeProject />
    <Footer />
    </>
  );
};
