import { Navbar } from "../navbar/navbar";

export const Hero = () => {
  return (
    <section className="relative w-full min-h-screen">
      <div className="flex items-center justify-center">
        <Navbar />
      </div>
      <div className="relative w-full flex-1 flex items-end md:items-center">
        <img
          src="https://minio.bocchikitsunei.com/bocchikitsuneiwebportfolio/Nei_PR_R_MH_CutOut.png"
          alt="Profile photo of BocchiKitsunei"
          className="absolute inset-0 w-full h-full object-cover object-right"
        />
        <div className="relative z-10 md:px-12 max-w-3xl">
          <h1 className="text-4xl md:text-7xl lg:text-8xl font-bold text-purple-500 leading-tight">
            BocchiKitsunei
          </h1>
          <p className="text-white md:text-black pb-3 text-lg md:text-2xl font-semibold">
            Nithit Lertcharoensombat{" "}
            <span className="font-extrabold text-purple-500">|</span> Software
            Engineer
          </p>
        </div>
      </div>
    </section>
  );
};
