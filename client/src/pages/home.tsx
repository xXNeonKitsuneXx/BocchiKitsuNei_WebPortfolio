import '../../app/globals.css';
import { Navbar } from '@/components/navbar/navbar';
import { Hero } from '@/components/homepage/hero';
import { About } from '@/components/homepage/about';
import { Skill } from '@/components/homepage/skill';
import { Project } from '@/components/homepage/project';

export const Home = () => {

  return (
    <>
    <Navbar />
    <Hero />
    <About />
    <Skill />
    <Project />
    <text>Kuy</text>
    <img className='w-20 mt-4' src='http://199.241.138.79:9000/bocchikitsuneiwebportfolio/ToDoList.png' />

    </>
  );
};
