import useSWR from "swr";
import { useState } from "react";
import { FaGithub } from "react-icons/fa";

interface Project {
  project_id: number;
  project_picture: string;
  project_name: string;
  project_tag: string;
  project_description: string;
  project_github_url: string;
  project_language: string;
}

const fetcher = (url: string): Promise<Project[]> =>
  fetch(url).then((res) => res.json());

const ProjectModal = ({ project, onClose }: { project: Project; onClose: () => void }) => {
  return (
    <div
      id="container"
      onClick={onClose}
      className="fixed inset-0 bg-[black] bg-opacity-25 backdrop-blur-sm flex justify-center items-center z-50"
    >
      <section className="flex justify-center">
        <div>
          <div>
            <div className="bg-white rounded-lg shadow-xl w-[375px] h-[590px] lg:hidden" data-aos="fade-up">
              <img
                className="rounded-t-lg w-[400px] h-[220px]"
                src={project.project_picture}
                alt={project.project_name}
              />
              <div className="p-5 w-[375px] h-[480px]">
                <div className="items-center">
                  <h5 className="mb-2 text-2xl font-bold tracking-tight">
                    <div className="text-purple-500 flex items-center">
                      {project.project_name}{" "}
                      <span className="text-base ml-2 bg-orange-400 rounded-xl py-1 px-2 text-white font-semibold">
                        {project.project_tag}
                      </span>
                    </div>
                  </h5>
                  <p className="mb-3 font-normal text-black overflow-y-scroll h-[150px]">
                    {project.project_description}
                  </p>
                  <h5 className="mb-2 text-2xl font-bold tracking-tight text-black">
                    Language
                  </h5>
                  <span className="mb-4 text-xl">
                    - {project.project_language}
                  </span>
                  <div className="flex flex-row mt-4">
                    {project.project_github_url && (
                      <a
                        href={project.project_github_url}
                        target="_blank"
                        className="inline-flex items-center justify-center mr-4 px-4 py-2 text-xl font-semibold text-center text-white rounded-lg bg-purple-500 hover:bg-purple-800"
                      >
                        <span className="mr-2 text-2xl">
                          <FaGithub />
                        </span>{" "}
                        Github
                      </a>
                    )}
                  </div>
                </div>
              </div>
            </div>

            <div className="hidden lg:flex bg-white rounded-lg shadow-xl flex-row h-[500px] w-[1100px]" data-aos="fade-up">
              <img
                className="object-cover w-[500px] h-[500px]"
                src={project.project_picture}
                alt={project.project_name}
              />
              <div className="flex flex-col p-8 leading-normal w-[600px] h-[500px]">
                <h5 className="mb-2 text-2xl font-bold tracking-tight">
                  <div className="text-purple-500 flex items-center">
                    {project.project_name}{" "}
                    <span className="text-base ml-2 bg-orange-400 rounded-xl py-1 px-2 text-white font-semibold">
                      {project.project_tag}
                    </span>
                  </div>
                </h5>
                <p className="mb-3 text-xl font-normal text-black overflow-y-scroll h-[275px]">
                  {project.project_description}
                </p>
                <h5 className="mb-2 text-2xl font-bold tracking-tight text-black">
                  Language
                </h5>
                <span className="mb-4 text-xl">
                  - {project.project_language}
                </span>

                <div className="flex flex-row">
                  {project.project_github_url && (
                    <a
                      href={project.project_github_url}
                      target="_blank"
                      className="inline-flex items-center justify-center mr-4 px-4 py-2 text-xl font-semibold text-center text-white rounded-lg bg-purple-500 hover:bg-purple-800"
                    >
                      <span className="mr-2 text-2xl">
                        <FaGithub />
                      </span>{" "}
                      Github
                    </a>
                  )}
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
};

export const Project = () => {
  const { data, error, isLoading } = useSWR<Project[]>("/api/GetProjectsFirstFour", fetcher);
  const [selectedProject, setSelectedProject] = useState<Project | null>(null);

  if (error) return <div>failed to load</div>;
  if (isLoading) return <div>loading...</div>;
  if (!data) return <div>loading...</div>;

  const handleCardClick = (project: Project) => {
    setSelectedProject(project);
  };

  const handleCloseModal = () => {
    setSelectedProject(null);
  };

  return (
    <section className="w-full py-8 md:py-10 lg:py-12 bg-white" data-aos="fade-up">
      <div className="flex justify-center">
        <div className="items-center font-bold" data-aos="fade-up">
          <p className="text-4xl md:text-6xl items-center font-bold">Projects</p>
          <hr className="mt-4 h-2 bg-purple-500 border-none w-full" />
        </div>
      </div>

      <div className="py-8 px-4 flex justify-center">
        <div className="grid gap-12 mb-8 lg:grid-cols-2">
          {data.map((project) => (
            <div
              key={project.project_id}
              className="bg-gray-50 rounded-lg shadow-xl w-[400px] h-[500px] hover:!scale-105 duration-500 hover:bg-gray-100 cursor-pointer"
              data-aos="flip-right"
              data-aos-duration="500"
              onClick={() => handleCardClick(project)}
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
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>

      {selectedProject && <ProjectModal project={selectedProject} onClose={handleCloseModal} />}
    </section>
  );
};
