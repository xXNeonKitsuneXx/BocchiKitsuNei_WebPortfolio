import { Navbar } from "../navbar/navbar";

export const Hero = () => {
  return (
    <section className="relative w-full min-h-screen">
      <Navbar />
      <div className="relative w-full h-[90vh] flex items-center">
        <img
          src="https://minio.bocchikitsunei.com/bocchikitsuneiwebportfolio/Nei_PR_R_MH.png"
          alt="Profile photo of BocchiKitsunei"
          className="absolute inset-0 w-full h-full object-cover object-right"
        />
        <div className="relative z-10 px-8 md:px-16 lg:px-24 max-w-2xl">
          <h1 className="text-5xl md:text-7xl lg:text-8xl font-bold text-purple-500 leading-tight">
            BocchiKitsunei
          </h1>
          <p className="text-black mt-4 text-xl md:text-2xl font-semibold">
            Nithit Lertcharoensombat{" "}
            <span className="font-extrabold text-purple-500">|</span> Software
            Developer
          </p>
        </div>
      </div>
    </section>
  );
};
