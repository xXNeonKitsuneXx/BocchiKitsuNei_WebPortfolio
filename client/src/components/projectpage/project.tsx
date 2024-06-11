import useSWR from "swr";
import { useState } from "react";
import { Dialog, DialogContent } from "@/components/ui/dialog";
import {
  Card,
  CardContent,
  CardDescription,
  CardTitle,
} from "@/components/ui/card";
import { FaGithub } from "react-icons/fa";
import { Navbar } from "../navbar/navbar";
import { LoadingSpinner } from "../loadingspinner/loadingspinner";

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

export const Project = () => {
  const { data, error, isLoading } = useSWR<Project[]>("/api/GetProjects", fetcher);

  const [isOpen, setIsOpen] = useState(false);
  const [selectedProject, setSelectedProject] = useState<Project | null>(null);

  if (error) return <div className="flex justify-center items-center mt-16 text-red-500">failed to load</div>;
  if (isLoading) return <div className="flex justify-center items-center mt-16"><LoadingSpinner/></div>;
  if (!data) return <div className="flex justify-center items-center mt-16"><LoadingSpinner /></div>;

  const handleCardClick = (project: Project) => {
    setSelectedProject(project);
    setIsOpen(true);
  };

  return (
    <section className="flex justify-center items-center">
      <Navbar />
      <div>
        <div className="mt-16 flex flex-col justify-center text-center">
          <h1
            className="mt-6 text-6xl font-bold text-black"
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
          <h4 className=" mt-4 px-5 text-lg md:text-2xl font-medium text-black border-l-[20px] border-purple-500 py-4">
            These are all the projects I have completed before. Most of them are
            open-source. You are welcome to use them in anyway you like.
          </h4>
        </div>
        <div
          className="w-full pt-6 bg-white"
          data-aos="fade-up"
        >
          <div className="px-4 flex justify-center">
            <div className="grid gap-8 mb-8 lg:grid-cols-3">
              {data.map((project) => (
                <Card
                  key={project.project_id}
                  onClick={() => handleCardClick(project)}
                  className="transition-all bg-gray-50 rounded-lg shadow-lg w-[350px] h-[500px] hover:!scale-105 duration-500 hover:bg-gray-100 cursor-pointer"
                  data-aos="flip-right"
                  data-aos-duration="500"
                >
                  <img
                    src={project.project_picture}
                    alt={project.project_name}
                    className="rounded-t-lg w-[350px] h-[260px]"
                  />
                  <CardContent className="px-4 pt-2 w-[350px] h-[200px]">
                    <CardTitle className="text-purple-500 mb-2 text-2xl font-bold tracking-tight">
                      {project.project_name}{" "}
                      <span className="text-base ml-2 bg-orange-400 rounded-xl py-1 px-2 text-white font-semibold">
                        {project.project_tag}
                      </span>{" "}
                    </CardTitle>
                    <CardDescription className="mb-3 font-normal text-black overflow-hidde line-clamp-8 h-[160px]">
                      {project.project_description}
                    </CardDescription>
                  </CardContent>
                </Card>
              ))}
              {/* //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// */}
              {selectedProject && (
                <Dialog open={isOpen} onOpenChange={setIsOpen}>
                  <DialogContent className="bg-white w-[375px] h-[600px] rounded-xl shadow-xl">
                    <img
                      src={selectedProject.project_picture}
                      alt={selectedProject.project_name}
                      className="rounded-t-lg w-[400px] h-[260px]"
                    />
                    <div className=" px-4 pt-0 mt-0 w-[375px] h-[480px]">
                      <h2 className="text-purple-500 mb-2 text-2xl font-bold tracking-tight">
                        {selectedProject.project_name}{" "}
                        <span className="text-base ml-2 bg-orange-400 rounded-xl py-1 px-2 text-white font-semibold">
                          {selectedProject.project_tag}
                        </span>
                      </h2>
                      <p className=" font-normal text-black overflow-y-scroll h-[150px]">
                        {selectedProject.project_description}
                      </p>
                      <h5 className="mb-2 mt-1 text-xl font-bold tracking-tight text-black">
                        Language
                      </h5>
                      <span className="mb-4 text-base">
                        - {selectedProject.project_language}
                      </span>
                      <div className="flex flex-row mt-2">
                        {selectedProject.project_github_url && (
                          <a
                            href={selectedProject.project_github_url}
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
                  </DialogContent>
                </Dialog>
              )}
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};
