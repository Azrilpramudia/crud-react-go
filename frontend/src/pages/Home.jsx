import Navbar from "../components/Navbar";

const Home = () => {
  return (
    <>
      <Navbar />
      <div className="flex justify-center items-center h-screen bg-blue-700">
        <h1 className="font-bold text-9xl">FULLSTACK</h1>
      </div>
    </>
  );
};

export default Home;