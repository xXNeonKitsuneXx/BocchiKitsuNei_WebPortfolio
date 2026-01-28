import { Navbar } from "../navbar/navbar";

export const Hero = () => {
  return (
    <section className="flex justify-center items-center">
      <Navbar />
      <div className="relative w-full h-[80vh] flex items-start justify-items-start" data-aos="zoom-in">
        <img src="https://minio.bocchikitsunei.com/bocchikitsuneiwebportfolio/Nei_PR_R_MH.png" alt="Background Image" className="absolute inset-0 w-full h-full object-cover" />
        <div className="relative z-10 text-purple-500">
          <h1 className="text-4xl md:text-6xl font-bold" data-aos="fade-up">BocchiKitsunei</h1>
          <p className="text-white mt-2 text-lg md:text-xl px-10 font-semibold text-center" data-aos="fade-up">Nithit Lertcharoensombat <span className="font-extrabold"> I </span> Sofeware Developer</p>
        </div>
      </div>
    </section>
  );
};
