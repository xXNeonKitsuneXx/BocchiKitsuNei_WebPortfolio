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
            <Button className="bg-black text-base font-bold hover:bg-gray-500">
              <a href="/">Home</a>
            </Button>
          </li>
          <li>
            <Button className="bg-black text-base font-bold hover:bg-gray-500">
              <a href="/project">Project</a>
            </Button>
          </li>
          <li>
            <Button className="bg-black text-base font-bold hover:bg-gray-500">
              <a href="/contact">Contact</a>
            </Button>
          </li>
        </ul>
      </div>
    </nav>
  );
};
