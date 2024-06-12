import { useState } from "react";
import { Card, CardContent } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import { Navbar } from "../navbar/navbar";
import { useToast } from "@/components/ui/use-toast";
import {
  FaEnvelope,
  FaFacebookSquare,
  FaGithub,
  FaInstagram,
  FaLinkedin,
} from "react-icons/fa";
import axios from "axios";

interface FormData {
  name: string;
  phone: string;
  email: string;
  message: string;
}

export const Contact = () => {
  const [formData, setFormData] = useState<FormData>({
    name: "",
    phone: "",
    email: "",
    message: "",
  });

  const [errors, setErrors] = useState<Partial<FormData>>({});
  const [copied, setCopied] = useState(false);
  const { toast } = useToast();

  const copyToClipboard = () => {
    const email = "BocchiKitsuNei@gmail.com";
    navigator.clipboard.writeText(email).then(() => {
      setCopied(true);
      toast({
        title: "Email copied to clipboard!",
      });
    });
  };

  const validateForm = () => {
    const newErrors: Partial<FormData> = {};
    if (!formData.name.trim()) {
      newErrors.name = "Name is required";
    }
    if (!formData.phone.trim()) {
      newErrors.phone = "Phone is required";
    }
    if (!formData.email.trim()) {
      newErrors.email = "Email is required";
    } else if (!/^\S+@\S+\.\S+$/.test(formData.email)) {
      newErrors.email = "Invalid email address";
    }
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const isValid = validateForm();
    if (isValid) {
      try {
        const response = await axios.post("/SendMail", {
          name: formData.name,
          email: formData.email,
          phone: formData.phone,
          msg: formData.message,
        });
  
        if (response.status === 200) {
          toast({
            title: "Email sent successfully!",
          });
          setFormData({ name: "", phone: "", email: "", message: "" });
        } else {
          toast({
            title: "Failed to send email.",
            description: "Please try again later.",
          });
        }
      } catch (error) {
        toast({
          title: "An error occurred.",
          description: "Please try again later.",
        });
      }
    }
  };
  
  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    setFormData({ ...formData, [e.target.id]: e.target.value });
  };

  return (
    <section className="flex justify-center items-center">
      <Navbar />
      <div className="mt-28 flex-col lg:flex-row flex justify-center items-center">
        <Card className="p-6 mb-4 lg:mr-8 w-96">
          <CardContent>
            <div className="space-y-8">
              <div className="space-y-2">
                <h2 className="text-3xl font-semibold">Contact me</h2>
                <div className="flex justify-center">
                  <hr className="mt-1 h-1.5 bg-purple-500 border-none w-full" />
                </div>
                <div className="py-1">
                  <div
                    onClick={copyToClipboard}
                    className="flex mt-1 hover:text-purple-500 hover:-translate-y-1 duration-300 cursor-pointer"
                  >
                    <FaEnvelope className="text-4xl md:text-5xl" />{" "}
                    <span className="pl-4 pt-1 md:pt-3 text-xl">
                      BocchiKitsuNei@gmail.com
                    </span>
                  </div>
                  {copied}

                  <div>
                    <a
                      href="https://www.facebook.com/xXNekoLordXx/"
                      target="_blank"
                      className="flex mt-5 hover:text-purple-500 hover:-translate-y-1 duration-300 cursor-pointer"
                    >
                      <FaFacebookSquare className="text-4xl md:text-5xl" />{" "}
                      <span className="pl-4 pt-1 md:pt-3 text-xl">
                        อยากรวย แต่ ไม่อยากเรียน เน
                      </span>
                    </a>
                  </div>

                  <div>
                    <a
                      href="https://www.instagram.com/kitsune_ne_cs/"
                      target="_blank"
                      className=" flex mt-5 hover:text-purple-500 hover:-translate-y-1 duration-300 cursor-pointer"
                    >
                      <FaInstagram className="text-4xl md:text-5xl" />{" "}
                      <span className="pl-4 pt-1 md:pt-3 text-xl">
                        @kitsune_ne_cs
                      </span>
                    </a>
                  </div>

                  <div>
                    <a
                      href="https://github.com/xXNeonKitsuneXx"
                      target="_blank"
                      className="flex mt-5 hover:text-purple-500 hover:-translate-y-1 duration-300 cursor-pointer"
                    >
                      <FaGithub className="text-4xl md:text-5xl" />{" "}
                      <span className="pl-4 pt-1 md:pt-3 text-xl">
                        xXNeonKitsuneXx
                      </span>
                    </a>
                  </div>

                  <div>
                    <a
                      href="https://www.linkedin.com/in/nithit-lertcharoensombat-722855264/"
                      target="_blank"
                      className="flex mt-5 hover:text-purple-500 hover:-translate-y-1 duration-300 cursor-pointer"
                    >
                      <FaLinkedin className="text-4xl md:text-5xl " />{" "}
                      <span className="pl-4 pt-1 md:pt-3 text-xl">
                        Nithit Lertcharoensombat
                      </span>
                    </a>
                  </div>
                </div>
              </div>
            </div>
          </CardContent>
        </Card>
        {/* ///////////////////////////////////////////////////////////////////////////////////////////////////// */}
        <Card className="p-6 mb-6 md:w-[550px] w-96">
          <CardContent className="">
            <div className="space-y-8">
              <div className="space-y-2">
                <h2 className="text-3xl font-semibold">Get in touch</h2>
                <p className="text-zinc-500 dark:text-zinc-400">
                  Fill out the form below and we'll get back to you as soon as
                  possible.
                </p>
              </div>
              <form onSubmit={handleSubmit}>
                <div className="space-y-4">
                  <div className="space-y-2">
                    <Label className="font-bold" htmlFor="name">
                      Name<span className="text-purple-500">*</span>
                    </Label>
                    <Input
                      id="name"
                      placeholder="Enter your name*"
                      value={formData.name}
                      onChange={handleChange}
                    />
                    {errors.name && (
                      <p className="text-red-500 text-sm">{errors.name}</p>
                    )}
                  </div>
                  <div className="space-y-2">
                    <Label className="font-bold" htmlFor="phone">
                      Phone<span className="text-purple-500">*</span>
                    </Label>
                    <Input
                      id="phone"
                      placeholder="Enter your phone*"
                      value={formData.phone}
                      onChange={handleChange}
                    />
                    {errors.phone && (
                      <p className="text-red-500 text-sm">{errors.phone}</p>
                    )}
                  </div>
                  <div className="space-y-2">
                    <Label className="font-bold" htmlFor="email">
                      Email<span className="text-purple-500">*</span>
                    </Label>
                    <Input
                      id="email"
                      placeholder="Enter your email*"
                      type="email"
                      value={formData.email}
                      onChange={handleChange}
                    />
                    {errors.email && (
                      <p className="text-red-500 text-sm">{errors.email}</p>
                    )}
                  </div>
                  <div className="space-y-2">
                    <Label className="font-bold" htmlFor="message">
                      Message
                    </Label>
                    <Textarea
                      id="message"
                      placeholder="Enter your message"
                      className="min-h-[100px]"
                      value={formData.message}
                      onChange={handleChange}
                    />
                  </div>
                  <Button
                    type="submit"
                    className="bg-purple-500 text-white
                    hover:bg-purple-800"
                  >
                    Submit
                  </Button>
                </div>
              </form>
            </div>
          </CardContent>
        </Card>
      </div>
    </section>
  );
};
