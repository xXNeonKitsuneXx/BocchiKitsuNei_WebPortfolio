import { Navbar } from "../navbar/navbar";

export const Hero = () => {
  return (
    <section className="relative w-full h-auto md:h-screen overflow-hidden">
      <div className="flex items-center justify-center"> 
        <Navbar />
      </div>
      <div className="relative w-full h-[634px] md:h-[934px] flex items-end justify-center md:justify-normal md:items-center overflow-hidden">
        <img
          src="https://minio.bocchikitsunei.com/bocchikitsuneiwebportfolio/Nei_Hero.jpg"
          alt="Profile photo of BocchiKitsunei"
          className="absolute inset-0 w-full h-[634px] md:h-[934px] object-cover object-right"
        />
        <div className="relative z-10 px-4 md:px-12 max-w-3xl">
          <h1 className="text-4xl md:text-7xl lg:text-8xl font-bold text-purple-500 leading-tight">
            Nithit Lertcharoensombat
          </h1>
          <p className="text-white md:text-black pb-3 text-lg md:text-2xl font-semibold">
            Software Engineer
          </p>
        </div>
      </div>
    </section>
  );
};
