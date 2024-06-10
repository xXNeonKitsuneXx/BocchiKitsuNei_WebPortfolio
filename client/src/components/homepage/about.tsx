import { Button } from "../ui/button";

export const About = () => {
  const handleDownloadResume = () => {
    window.location.href = "http://199.241.138.79:9000/bocchikitsuneiwebportfolio/Resume_Nithit_Lertcharoensombat.pdf";
  };

  return (
    <section className="w-full py-8 md:py-12 bg-white">
      <div className="container space-y-12 px-4 md:px-6">
        <div className="mx-auto grid max-w-5xl items-center gap-6 lg:grid-cols-2 lg:gap-12">
          <div className="flex flex-col justify-center space-y-4">
            <ul className="grid gap-6">
              <li>
                <div className="grid gap-1">
                  <h1 className="text-3xl md:text-5xl font-bold" data-aos="zoom-in-up">
                    Nithit Lertcharoensombat
                  </h1>
                  <p className="mt-2 text-lg md:text-2xl text-gray-500" data-aos="zoom-in-up">
                    A highly motivated Second-year computer science student
                    seeking an internship in a Software Developer. Seeking an
                    opportunity to apply my knowledge, skills and gain hands-on
                    experience in the field.
                  </p>
                </div>
                <Button
                  className="mt-2 md:mt-8 bg-purple-500 hover:bg-purple-800"
                  data-aos="zoom-in-up"
                  onClick={handleDownloadResume}
                >
                  Download Resume
                </Button>
              </li>
            </ul>
          </div>
          <img
            src="http://199.241.138.79:9000/bocchikitsuneiwebportfolio/Picture_Nei_Swimmer_3.jpg"
            alt="Picture_Nei_Swimmer"
            className="w-[550px] h-[310px] md:w-[400px] md:h-[] mx-auto xl:aspect-square aspect-video overflow-hidden rounded-xl object-cover sm:w-full lg:order-last"
            data-aos="flip-left"
            data-aos-duration="400"
          />
        </div>
      </div>
    </section>
  );
};
