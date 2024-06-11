import useSWR from "swr";
import { Navbar } from "../navbar/navbar";


interface Project {
  project_id: number;
  project_picture: string;
  project_name: string;
  project_tag: string;
  project_description: string;
  project_github_url: string;
}

const fetcher = (url: string): Promise<Project[]> =>
  fetch(url).then((res) => res.json());

export const Project = () => {
  const { data, error, isLoading } = useSWR<Project[]>(
    "/api/GetProjects",
    fetcher
  );

  if (error) return <div>failed to load</div>;
  if (isLoading) return <div>loading...</div>;
  if (!data) return <div>loading...</div>;

  return (
    <div className="flex justify-center items-center">
      <Navbar />
      <section>
        <div className="mt-16 flex flex-col justify-center text-center">
          <h1
            className="mt-[20px] text-6xl lg:text-8xl font-bold text-black"
            data-aos="fade-up"
            data-aos-duration="1000"
          >
            Project
          </h1>
          <div className="flex justify-center px-8">
            <hr
              className="mt-5 h-2 bg-black border-none w-full"
              data-aos="fade-up"
              data-aos-duration="1000"
            />
          </div>
        </div>
        <div
          className="px-12 flex justify-center"
          data-aos="fade-up"
          data-aos-duration="1000"
        >
          <h4 className=" mt-4 px-5 text-2xl md:text-4xl font-medium text-black border-l-[20px] border-purple-500 py-4">
            These are all the projects I have completed before. Most of them are
            open-source. You are welcome to use them in anyway you like.
          </h4>
        </div>
        <div
          className="w-full py-4 md:py-10 lg:py-12 bg-white"
          data-aos="fade-up"
        >
          <div className="px-4 flex justify-center">
            <div className="grid gap-12 mb-8 lg:grid-cols-3">
              {data.map((project) => (
                <div
                  key={project.project_id}
                  className="bg-gray-50 rounded-lg shadow-xl w-[400px] h-[500px] hover:!scale-105 duration-500"
                  data-aos="flip-right"
                  data-aos-duration="500"
                >
                  <img
                    className="rounded-t-lg w-[400px] h-[220px]"
                    src={project.project_picture}
                    alt={project.project_name}
                  />
                  <div className="p-5 w-[400px] h-[280px]">
                    <div className="items-center">
                      <h5 className="mb-2 text-2xl font-bold tracking-tight text-purple-500">
                        <div className="flex items-center">
                          {project.project_name}{" "}
                          <span className="text-base ml-2 bg-orange-400 rounded-xl py-1 px-2 text-white font-semibold">
                            {project.project_tag}
                          </span>
                        </div>
                      </h5>
                      <p className="mb-3 font-normal text-black overflow-hidden h-[190px]">
                        {project.project_description}
                      </p>
                      <a
                        href={project.project_github_url}
                        className="text-blue-500 hover:underline"
                        target="_blank"
                        rel="noopener noreferrer"
                      >
                        View on GitHub
                      </a>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </section>
    </div>
  );
};
