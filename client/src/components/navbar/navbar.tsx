import { Button } from "../ui/button";

export const Navbar = () => {
  return (
    <nav className="fix w-full bg-black p-1 rounded-full flex justify-between items-center">
      {/* <div className="text-cyan-400 text-2xl">Icon</div> */}
      <div className="flex-1 flex justify-center">
        <ul className="flex space-x-16 text-white">
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
        {/* <div className="text-lg">
            <span className="text-cyan-400">EN</span>
            <span className="text-gray-400"> | </span>
            <span className="text-cyan-400">TH</span>
          </div> */}
      </div>
    </nav>
  );
};
