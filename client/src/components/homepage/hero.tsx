import { Navbar } from "../navbar/navbar";

export const Hero = () => {
  return (
    <section className="relative w-full min-h-screen">
      <Navbar />
      <div className="relative w-full h-[90vh] flex items-end md:items-center md:pb-0">
        <img
          src="https://minio.bocchikitsunei.com/bocchikitsuneiwebportfolio/Nei_PR_R_MH_CutOut.png"
          alt="Profile photo of BocchiKitsunei"
          className="absolute inset-0 w-full h-full object-cover object-right"
        />
        <div className="relative z-10 px-8 max-w-3xl">
          <h1 className="text-5xl md:text-7xl lg:text-8xl font-bold text-purple-500 leading-tight">
            BocchiKitsunei
          </h1>
          <p className="text-white mt-4 md:text-black text-xl md:text-2xl font-semibold">
            Nithit Lertcharoensombat{" "}
            <span className="font-extrabold text-purple-500">|</span> Software
            Engineer
          </p>
        </div>
      </div>
    </section>
  );
};
