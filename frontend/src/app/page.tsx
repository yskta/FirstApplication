"use client";
import { useState, useEffect } from 'react';

export default function Home() {
  const [projects, setProjects] = useState([]);

  useEffect(() => {
    // ここでバックエンドからプロジェクトデータを取得します
    // 例: fetchProjects().then(data => setProjects(data));
  }, []);

  return (
    <main className="min-h-screen p-8">
      <h1 className="text-3xl font-bold mb-6">TaskMaster</h1>
      
      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-4">Projects</h2>
        <ul className="space-y-2">
          {projects.map(project => (
            <li key={project.id} className="bg-gray-100 p-4 rounded">
              <h3 className="text-xl font-medium">{project.name}</h3>
              <p className="text-gray-600">{project.description}</p>
            </li>
          ))}
        </ul>
      </section>

      <section>
        <h2 className="text-2xl font-semibold mb-4">Create New Project</h2>
        <form className="space-y-4">
          <div>
            <label htmlFor="name" className="block mb-1">Name</label>
            <input type="text" id="name" className="w-full border p-2 rounded" />
          </div>
          <div>
            <label htmlFor="description" className="block mb-1">Description</label>
            <textarea id="description" className="w-full border p-2 rounded" rows="3"></textarea>
          </div>
          <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600">
            Create Project
          </button>
        </form>
      </section>
    </main>
  );
}
