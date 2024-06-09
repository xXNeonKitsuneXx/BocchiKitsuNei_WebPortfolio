import useSWR from "swr";

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
    "/api/GetProjectsFirstFour",
    fetcher
  );

  if (error) return <div>failed to load</div>;
  if (isLoading) return <div>loading...</div>;
  if (!data) return <div>loading...</div>;

  return (
    <section
      className="w-full py-8 md:py-10 lg:py-12 bg-white"
      data-aos="fade-up"
    >
      <div className="flex justify-center">
        <div className="items-center font-bold" data-aos="fade-up">
          <p className="text-4xl md:text-6xl items-center font-bold">
            Projects
          </p>
          <hr className="mt-4 h-2 bg-purple-500 border-none w-full" />
        </div>
      </div>

      <div className="py-8 px-4 flex justify-center">
        <div className="grid gap-12 mb-8 lg:grid-cols-2">
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
                  <h5 className="mb-2 text-2xl font-bold tracking-tight text-[#7E38B7]">
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
    </section>
  );
};
