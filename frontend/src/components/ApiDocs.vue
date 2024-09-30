<template>
    <div class="api-docs">
        <div class="api-section" v-for="api in apis" :key="api.url">
            <div class="api-header" @click="toggleEndpoint(api)">
                <span :class="api.method">{{ api.method }}</span> <span>{{ api.url }}</span>
            </div>

            <transition name="fade">
                <div v-if="api.open" class="api-content">
                    <p>{{ api.description }}</p>

                    <!-- For GET method -->
                    <div v-if="api.method === 'GET' && api.url === '/api/books'">
                        <button @click="fetchBooks">Try it</button>
                        <div v-if="books.length > 0">
                            <h4>Response:</h4>
                            <div class="json-box">
                                <pre><code>{{ formatJson(books) }}</code></pre>
                            </div>
                        </div>
                    </div>

                    <!-- For POST method -->
                    <div v-if="api.method === 'POST' && api.url === '/api/books'">
                        <input v-model="createdBook.title" placeholder="Title" />
                        <input v-model="createdBook.author" placeholder="Author" />
                        <input v-model="createdBook.genre" placeholder="Genre" />
                        <input v-model="createdBook.desc" placeholder="Description" />
                        <input v-model="createdBook.isbn" placeholder="Isbn" />
                        <input type="file" @change="handleImageUpload($event, 'create')" />
                        <input type="date" v-model="createdBook.published" placeholder="Published" />
                        <input v-model="createdBook.publisher" placeholder="Publisher" />
                        <button @click="createBookFunc">Create Book</button>
                        <div v-if="responseMessage">
                            <h4>Response:</h4>
                            <div class="json-box">
                                <pre><code>{{ formatJson(responseMessage) }}</code></pre>
                            </div>
                        </div>
                    </div>

                    <!-- For PATCH method -->
                    <div v-if="api.method === 'PATCH' && api.url === '/api/books/:id'">
                        <input type="number" v-model="updateBookId" placeholder="Book ID"/>
                        <input v-model="updateBook.title" placeholder="Title" />
                        <input v-model="updateBook.author" placeholder="Author" />
                        <input v-model="updateBook.genre" placeholder="Genre" />
                        <input v-model="updateBook.desc" placeholder="Description" />
                        <input v-model="updateBook.isbn" placeholder="Isbn" />
                        <input type="file" @change="handleImageUpload($event, 'update')" />
                        <input type="date" v-model="updateBook.published" placeholder="Published" />
                        <input v-model="updateBook.publisher" placeholder="Publisher" />
                        <button @click="updateBookFunc">Update Book</button>

                        <div v-if="responseMessage">
                            <h4>Response:</h4>
                            <div class="json-box">
                                <pre><code>{{ formatJson(responseMessage) }}</code></pre>
                            </div>
                        </div>
                    </div>

                    <!-- For DELETE method -->
                    <div v-if="api.method === 'DELETE' && api.url === '/api/books/:id'">
                        <input v-model="deleteBookId" placeholder="Book ID" />
                        <button @click="deleteBookFunc">Delete Book</button>
                        <div v-if="responseMessage">
                            <h4>Response:</h4>
                            <div class="json-box">
                                <pre><code>{{ formatJson(responseMessage) }}</code></pre>
                            </div>
                        </div>
                    </div>

                    <!-- For Reset method -->
                    <div v-if="api.method === 'POST' && api.url === '/api/books/reset'">
                        <button @click="resetBooks">Reset Books</button>
                        <div v-if="responseMessage">
                            <h4>Response:</h4>
                            <div class="json-box">
                                <pre><code>{{ formatJson(responseMessage) }}</code></pre>
                            </div>
                        </div>
                    </div>
                </div>
            </transition>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            // Book data
            apiUrl: process.env.VUE_APP_API_URL,
            books: [],
            responseMessage: null,
            createdBook: {},
            updateBook: {},
            updateBookId: '',
            deleteBookId: '',
            imageFile: null,

            // API endpoints for documentation
            apis: [
                { method: 'GET', url: '/api/books', description: 'Fetch all books', open: false },
                { method: 'POST', url: '/api/books', description: 'Create a new book', open: false },
                { method: 'PATCH', url: '/api/books/:id', description: 'Update a book by ID', open: false },
                { method: 'DELETE', url: '/api/books/:id', description: 'Delete a book by ID', open: false },
                { method: 'POST', url: '/api/books/reset', description: 'Reset the book list', open: false }
            ]
        };
    },
    created() {
        console.log('API URL:', this.apiUrl);
    },
    methods: {
        // Toggle endpoint details
        toggleEndpoint(api) {
            api.open = !api.open;
            this.responseMessage = null;
        },

        handleImageUpload(event, type) {
            const file = event.target.files[0];
            if (type === 'create') {
                this.imageFile = file; // Assign the file to the create imageFile
            } else {
                this.updateBook.imageFile = file; // Assign the file to the update imageFile
            }
        },

        // Fetch all books
        fetchBooks() {
            axios.get(`${this.apiUrl}/api/books`)
                .then(response => {
                    this.books = response.data;
                })
                .catch(error => {
                    console.error('Error fetching books:', error);
                });
        },

        // Create a new book
        createBookFunc() {
            const formData = new FormData();
            formData.append('title', this.createdBook.title || "");
            formData.append('author', this.createdBook.author || "");
            formData.append('genre', this.createdBook.genre || "");
            formData.append('desc', this.createdBook.desc || "");
            formData.append('isbn', this.createdBook.isbn || "");
            formData.append('published', this.createdBook.published || "");
            formData.append('publisher', this.createdBook.publisher || "");

            if (this.imageFile) {
                formData.append('image', this.imageFile);
            }

            axios.post(`${this.apiUrl}/api/books`, formData, {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
            })
                .then(response => {
                    console.log(response.data)
                    this.createdBook = response.data; // Store created book details
                    this.responseMessage = response.data;
                    this.createdBook = { title: '', author: '', genre: '', desc: '', isbn: '', published: '', publisher: '' }; // Reset input fields
                })
                .catch(error => {
                    this.responseMessage = error.message;
                    this.createdBook = null; // Clear created book data
                });
        },

        updateBookFunc() {
            const formData = new FormData();
            // Only add fields that have values
            if (this.updateBook.title) {
                formData.append('title', this.updateBook.title);
            }
            if (this.updateBook.author) {
                formData.append('author', this.updateBook.author);
            }
            if (this.updateBook.genre) {
                formData.append('genre', this.updateBook.genre);
            }
            if (this.updateBook.desc) {
                formData.append('desc', this.updateBook.desc);
            }
            if (this.updateBook.isbn) {
                formData.append('isbn', this.updateBook.isbn);
            }
            if (this.updateBook.published) {
                formData.append('published', this.updateBook.published);
            }
            if (this.updateBook.publisher) {
                formData.append('publisher', this.updateBook.publisher);
            }

            // Handle the image file only if it exists
            if (this.updateBook.imageFile) {
                formData.append('image', this.updateBook.imageFile);
            }

            console.log("formdata", formData)
            axios.patch(`${this.apiUrl}/api/books/${this.updateBookId}`, formData, {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
            })
                .then(response => {
                    const index = this.books.findIndex(book => book.id === this.updateBookId);
                    if (index !== -1) {
                        this.books[index] = response.data;
                    }
                    this.responseMessage = response.data
                    this.updateBookId = '';
                    this.updateBook = {};
                })
                .catch(error => {
                    this.responseMessage = error.message;
                    this.updateBook = null;
                    console.error('Error updating book:', error);
                });
        },

        // Delete a book by ID
        deleteBookFunc() {
            axios.delete(`${this.apiUrl}/api/books/${this.deleteBookId}`)
                .then(() => {
                    this.books = this.books.filter(book => book.id !== this.deleteBookId);
                    this.deleteBookId = '';
                    this.responseMessage = { "message": "Deleted successfully" }
                })
                .catch(error => {
                    this.responseMessage = error.message;
                    console.error('Error deleting book:', error);
                });
        },

        // Reset the book list to the original state
        resetBooks() {
            axios.post(`${this.apiUrl}/api/books/reset`)
                .then(response => {
                    this.books = response.data;
                    this.responseMessage = response.data;
                })
                .catch(error => {
                    this.responseMessage = error.message;
                    console.error('Error resetting books:', error);
                });
        },
        // json format for display
        formatJson(data) {
            return JSON.stringify(data, null, 2); // Indent with 2 spaces
        }
    }
};
</script>

<style scoped>
.api-docs {
    text-align: left;
    margin: 0 auto;
    width: 80%;
}

.api-section {
    border: 1px solid #eaeaea;
    margin: 10px 0;
    padding: 15px;
    border-radius: 8px;
    background-color: #f7f7f7;
}

.api-header {
    display: flex;
    justify-content: space-between;
    cursor: pointer;
    padding: 10px;
    background-color: #f1f1f1;
    border-radius: 5px;
    font-weight: bold;
    align-items: center;
}

.api-header span {
    padding: 5px 10px;
    border-radius: 5px;
}

.api-header .GET {
    background-color: #28a745;
    color: white;
}

.api-header .POST {
    background-color: #007bff;
    color: white;
}

.api-header .PATCH {
    background-color: #ffc107;
    color: white;
}

.api-header .DELETE {
    background-color: #dc3545;
    color: white;
}

.api-content {
    margin-top: 10px;
    padding: 15px;
    border-top: 1px solid #eaeaea;
}

button {
    background-color: #007bff;
    color: white;
    border: none;
    padding: 8px 16px;
    cursor: pointer;
    border-radius: 5px;
}

button:hover {
    background-color: #0056b3;
}

input {
    padding: 8px;
    margin: 5px 0;
    border: 1px solid #ccc;
    border-radius: 5px;
    width: calc(100% - 20px);
}

ul {
    list-style-type: none;
    padding: 0;
}

li {
    background-color: #f7f7f7;
    padding: 10px;
    margin: 5px 0;
    border-radius: 5px;
}

.json-box {
    background-color: #f7f7f7;
    /* Light background color */
    border: 1px solid #ccc;
    /* Border around the box */
    border-radius: 5px;
    /* Rounded corners */
    padding: 10px;
    /* Padding inside the box */
    overflow: auto;
    /* Enable scrolling if content is too long */
    max-height: 300px;
    /* Limit height to avoid overflow */
}
</style>