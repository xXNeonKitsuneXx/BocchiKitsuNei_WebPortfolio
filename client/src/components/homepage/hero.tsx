import { Navbar } from "../navbar/navbar";

export const Hero = () => {
  return (
    <div className="flex justify-center items-center">
      <Navbar />
      <section className="relative w-full h-[80vh] flex items-center justify-center" data-aos="zoom-in">
        <img src="http://199.241.138.79:9000/bocchikitsuneiwebportfolio/Nei.jpg" alt="Background Image" className="absolute inset-0 w-full h-full object-cover" />
        <div className="relative z-10 text-center text-white">
          <h1 className="text-4xl md:text-6xl font-bold" data-aos="fade-up">BocchiKitsunei</h1>
          <p className="mt-2 text-lg md:text-xl px-10 font-semibold text-center" data-aos="fade-up">Nithit Lertcharoensombat <span className="font-extrabold"> I </span> Sofeware Developer</p>
        </div>
      </section>
    </div>
  );
};
