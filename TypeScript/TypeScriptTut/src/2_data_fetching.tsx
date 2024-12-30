import React from 'react'

type PostType = {
  id: number;
  title: string;
  desc: string;
}

const _2_data_fetching = async () => {
  const GetPosts = async () => {
      const response = await fetch(
          "https://jsonplaceholder.typicode.com/posts"
      );

      if (response.ok) {
          return response.json();
      }
      throw new Error("Something went wrong");
  };
  const data: PostType[] = await GetPosts();
  return (
      <div>
          {data.map((post: PostType) => (
              <div key={post.id}>
                  <h1>{post.title}</h1>
                  <p>{post.desc}</p>
              </div>
          ))}
      </div>
  );
}

export default _2_data_fetching
