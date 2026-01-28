import { Button } from "../ui/button";

export const Navbar = () => {
  return (
    <nav
      className="fixed top-0 p-2 flex justify-center mt-6 items-center bg-black z-20"
      style={{ borderRadius: '9999px', width: 'max-content' }}
      data-aos="fade-down"
    >
      <div className="flex justify-center px-4">
        <ul className="flex space-x-8 text-white">
          <li>
            <Button className="text-white font-medium hover:text-purple-400 transition-colors">
              <a href="/">Home</a>
            </Button>
          </li>
          <li>
            <Button className="text-white font-medium hover:text-purple-400 transition-colors">
              <a href="/project">Project</a>
            </Button>
          </li>
          <li>
            <Button className="text-white font-medium hover:text-purple-400 transition-colors">
              <a href="/contact">Contact</a>
            </Button>
          </li>
        </ul>
      </div>
    </nav>
  );
};
