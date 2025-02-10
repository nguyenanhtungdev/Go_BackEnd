import { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [tasks, setTasks] = useState([]);
  const [showForm, setShowForm] = useState(false);
  const [selectedIndex, setSelectedIndex] = useState(null);
  const [message, setMessage] = useState("");
  const [messageType, setMessageType] = useState("");
  const [formData, setFormData] = useState({
    productID: "",
    productName: "",
    price: "",
    description: "",
  });

  const handleAction = () => {
    setMessage("Thành công! Hành động đã được thực hiện.");
    setMessageType("success");
    setTimeout(() => {
      setMessage("");
    }, 3000);
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleOnClick = (row, index) => {
    setFormData({
      productID: row.productID,
      productName: row.productName,
      price: row.price,
      description: row.description,
    });
    setShowForm(true);
    setSelectedIndex(index);
  };

  useEffect(() => {
    fetch("http://localhost:8020/products")
      .then((response) => response.json())
      .then((data) => {
        setTasks(data);
      })
      .catch((error) => console.error("Error fetching tasks:", error));
  }, []);

  const fetchSortedProducts = async (field, order) => {
    fetch(
      `http://localhost:8020/products/sorted?sortField=${field}&sortOrder=${order}`
    )
      .then((response) => response.json())
      .then((data) => {
        setTasks(data);
      })
      .catch((error) => console.error("Error fetching tasks:", error));
  };

  const handleInsert = async () => {
    try {
      const response = await fetch("http://localhost:8020/products/add", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ...formData,
          price: parseInt(formData.price, 10),
        }),
      });

      if (response.ok) {
        handleAction();
        setFormData(false);

        fetch("http://localhost:8020/products")
          .then((response) => response.json())
          .then((data) => {
            setTasks(null);
            setTasks(data);
          })
          .catch((error) => console.error("Error fetching tasks:", error));
      } else {
        const errorText = await response.text();
        console.error("Lỗi từ server:", errorText);
        alert("Lỗi khi thêm mới sản phẩm: " + errorText);
      }
    } catch (error) {
      console.error("Lỗi kết nối:", error);
      alert("Không thể kết nối tới server");
    }
  };

  const handleUpdate = async () => {
    try {
      const response = await fetch("http://localhost:8020/products/update", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ...formData,
          price: parseInt(formData.price, 10),
          productID: parseInt(formData.productID),
        }),
      });

      if (response.ok) {
        handleAction();
        setShowForm(false);

        fetch("http://localhost:8020/products")
          .then((response) => response.json())
          .then((data) => {
            setTasks(null);
            setTasks(data);
          })
          .catch((error) => console.error("Error fetching tasks:", error));
      } else {
        const errorText = await response.text();
        console.error("Lỗi từ server:", errorText);
        alert("Lỗi khi cập nhật sản phẩm: " + errorText);
      }
    } catch (error) {
      console.error("Lỗi kết nối:", error);
      alert("Không thể kết nối tới server");
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>Danh sách sản phẩm</h1>
        <table className="product-table">
          <thead>
            <tr>
              <th>Mã sản phẩm</th>
              <th>Tên sản phẩm</th>
              <th>Giá bán</th>
              <th>Mô tả</th>
            </tr>
          </thead>
          <tbody>
            {tasks.map((product, index) => (
              <tr
                key={product.productID}
                onClick={() => handleOnClick(product, index)}
                className={selectedIndex === index ? "selectedRow" : ""}
              >
                <td>{product.productID}</td>
                <td>{product.productName}</td>
                <td>{product.price}</td>
                <td>{product.description ? product.description : "Trống"}</td>
              </tr>
            ))}
          </tbody>
        </table>

        {showForm && (
          <div className="inputForm">
            <input
              type="text"
              name="productID"
              placeholder="Mã sản phẩm"
              value={formData.productID}
              onChange={handleChange}
            />
            <input
              type="text"
              name="productName"
              placeholder="Tên sản phẩm"
              value={formData.productName}
              onChange={handleChange}
            />
            <input
              type="text"
              name="price"
              placeholder="Giá bán"
              value={formData.price}
              onChange={handleChange}
            />
            <input
              type="text"
              name="description"
              placeholder="Mô tả"
              value={formData.description}
              onChange={handleChange}
            />
            <div className="listBtn">
              <button className="btn-Update" onClick={handleUpdate}>
                Cập nhật
              </button>
              <button className="btn-Update" onClick={handleInsert}>
                Thêm mới
              </button>
            </div>
          </div>
        )}

        {/* Hiển thị thông báo */}
        {message && <div className={`alert ${messageType}`}>{message}</div>}
        <div className="listBtn_1">
          <button className="btn-Update" onClick={() => setShowForm(!showForm)}>
            {!showForm ? "Quản lý sản phẩm" : "Hoàn tất"}
          </button>
          {!showForm && (
            <>
              <button
                className="btn-Update"
                onClick={() => fetchSortedProducts("price", "asc")}
              >
                Sắp xếp tăng
              </button>
              <button
                className="btn-Update"
                onClick={() => fetchSortedProducts("price", "desc")}
              >
                Sắp xếp giảm
              </button>
            </>
          )}
        </div>
      </header>
    </div>
  );
}

export default App;
